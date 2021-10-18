<template>
  <div>
    <h3>用户列表页面</h3>
    <a-card>
      <a-row>
        <a-col :span="6">
          <a-input-search placeholder="输入用户名查找" enter="button" />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">新增</a-button>
        </a-col>
      </a-row>

      <a-table rowKey="username" :columns="columns" :pagination="paginationOption" :dataSource="userList" bordered>
          <span slot="role" slot-scope="role">{{role == 1 ? '管理员':'订阅者'}}</span>
          <template slot="action">
            <div class="actionSlot">
              <a-button type="primary">编辑</a-button>
              <a-button type="danger">删除</a-button>
            </div>
          </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '10%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    scopedSlots: { customRender: 'role' },
    align: 'center'
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center'
  }
]
export default {
  data () {
    return {
      paginationOption: {
        pageSizeOptions: ['2', '5', '10'],
        defaultCurrent: 1,
        defaultPageSize: 2,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        onChange: (current, pageSize) => {
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = pageSize
          this.getUserList()
        },
        onshowSizeChange: (current, size) => {
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = size
          this.getUserList()
        }
      },
      userList: [],
      columns
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    async getUserList () {
      const { data: res } = await this.$http.get('users', {
        params: {
          pageSize: this.paginationOption.defaultPageSize,
          pagenum: this.paginationOption.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.userList = res.data
      this.paginationOption.total = res.total
    }
  }
}
</script>

<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}
</style>
