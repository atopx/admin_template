<script lang="ts" setup>
import { reactive, ref, watch } from "vue"
import { getUserListApi, deleteUserApi, createUserApi, updateUserApi, disableUserApi } from "@/api/user"
import { UserListRequestData, UserInfo } from "@/api/user/types"
import { type FormInstance, ElMessage, ElMessageBox } from "element-plus"
import { Search, Refresh, CirclePlus, RefreshRight } from "@element-plus/icons-vue"
import { usePagination } from "@/hooks/usePagination"

defineOptions({
    name: "UserList"
})

const loading = ref<boolean>(false)

const { paginationData, handleCurrentChange, handleSizeChange } = usePagination()
const handleDelete = (row: UserInfo) => {
    ElMessageBox.confirm(`< ${row.username} >`, "删除用户", {
        confirmButtonText: "删除",
        cancelButtonText: "取消",
        type: "warning"
    }).then(() => {
        deleteUserApi(row.userId).then(() => {
            ElMessage.success("删除成功")
            handleRefresh()
        }).catch((err) => {
            ElMessage.error(err)
        })
    })
}

const handleDisable = (row: UserInfo) => {
    ElMessageBox.confirm(`< ${row.username} >`, "禁用用户", {
        confirmButtonText: "禁用",
        cancelButtonText: "取消",
        type: "warning"
    }).then(() => {
        disableUserApi(row.userId).then(() => {
            ElMessage.success("禁用成功")
            handleRefresh()
        }).catch((err) => {
            ElMessage.error(err)
        })
    })
}

const tableData = ref<UserInfo[]>([])
const searchFormRef = ref<FormInstance | null>(null)
const searchData = reactive({
    keyword: "",
    level: "",
    timeRange: ""
})

const handleCreate = () => {
    console.log("todo handleCreate")
}

const handleUpdate = (row: UserInfo) => {
    console.log("todo handleUpdate")
}

// 刷新列表
const handleRefresh = () => {
    loading.value = true

    const params: UserListRequestData = {
        pageInfo: {
            index: paginationData.currentPage,
            size: paginationData.pageSize,
            count: paginationData.total,
            disable: false
        },
        filter: {
            keyword: searchData.keyword,
            timeRange: {
                left: 0,
                right: 0
            }
        },
    }

    if (searchData.timeRange.length === 2) {
        params.filter.timeRange = {
            left: (searchData.timeRange[0] as unknown as Date).getTime() / 1000,
            right: (searchData.timeRange[1] as unknown as Date).getTime() / 1000 + 86400
        }
    }

    getUserListApi(params)
        .then((resp) => {
            paginationData.total = resp.data.pageInfo.count
            tableData.value = resp.data.list
        })
        .catch(() => {
            tableData.value = []
            console.log(2)
        })
        .finally(() => {
            loading.value = false
            console.log(3)
        })
}

// 搜索
const handleSearch = () => {
    if (paginationData.currentPage === 1) {
        handleRefresh()
    }
    paginationData.currentPage = 1
}

// 重置搜索条件
const resetSearch = () => {
    searchFormRef.value?.resetFields()
    if (paginationData.currentPage === 1) {
        handleRefresh()
    }
    paginationData.currentPage = 1
}

/** 监听分页参数的变化 */
watch([() => paginationData.currentPage, () => paginationData.pageSize], handleRefresh, { immediate: true })
</script>

<template>
    <div class="app-container">
        <el-card v-loading="loading" shadow="never" class="search-wrapper">
            <el-form ref="searchFormRef" :inline="true" :model="searchData">
                <el-form-item prop="keyword" label="关键词">
                    <el-input v-model="searchData.keyword" placeholder="关键词" />
                </el-form-item>
                <el-form-item prop="timeRange" label="时间">
                    <el-date-picker v-model="searchData.timeRange" type="daterange" range-separator="至"
                        start-placeholder="开始日期" end-placeholder="结束日期" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" :icon="Search" @click="handleSearch">查询</el-button>
                    <el-button :icon="Refresh" @click="resetSearch">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card v-loading="loading" shadow="never">
            <div class="toolbar-wrapper">
                <div>
                    <el-button type="primary" :icon="CirclePlus" @click="handleCreate()">新增用户</el-button>
                </div>
                <div>
                    <el-tooltip content="刷新表格">
                        <el-button type="primary" :icon="RefreshRight" circle @click="handleRefresh" />
                    </el-tooltip>
                </div>
            </div>
            <div class="table-wrapper">
                <el-table :data="tableData">
                    <el-table-column width="50" align="center" />
                    <el-table-column prop="userId" label="ID" width="50" align="center" />
                    <el-table-column prop="username" label="用户名" align="center" />
                    <el-table-column label="状态" align="center">
                        <template #default="scope">
                            <el-tag>{{ scope.row.status }}</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="createTime" label="创建时间" align="center" />
                    <el-table-column prop="updateTime" label="更新时间" align="center" />
                    <el-table-column fixed="right" label="操作" width="200" align="center">
                        <template #default="scope">
                            <el-button type="primary" text bg size="small" @click="handleUpdate(scope.row)">编辑</el-button>
                            <el-button type="warning" text bg size="small" @click="handleDisable(scope.row)">禁用</el-button>
                            <el-button type="danger" text bg size="small" @click="handleDelete(scope.row)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>

            <div class="pager-wrapper">
                <el-pagination background :layout="paginationData.layout" :page-sizes="paginationData.pageSizes"
                    :total="paginationData.total" :page-size="paginationData.pageSize"
                    :currentPage="paginationData.currentPage" @size-change="handleSizeChange"
                    @current-change="handleCurrentChange" />
            </div>
        </el-card>
    </div>
</template>

<style lang="scss" scoped>
.search-wrapper {
    margin-bottom: 20px;

    :deep(.el-card__body) {
        padding-bottom: 2px;
    }
}

.toolbar-wrapper {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
}

.table-wrapper {
    margin-bottom: 20px;
}

.pager-wrapper {
    display: flex;
    justify-content: flex-end;
}
</style>
