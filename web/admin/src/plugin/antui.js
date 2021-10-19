import Vue from 'vue'
import { Button, FormModel, Input, Icon, message, Layout, Menu, Card, Table, Row, Col, ConfigProvider } from 'ant-design-vue'

Vue.prototype.$message = message

message.config({
  top: '60px',
  duration: 2,
  maxCount: 3
})

Vue.use(Button)
Vue.use(Input)
Vue.use(FormModel)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
Vue.use(Card)
Vue.use(Table)
Vue.use(Row)
Vue.use(Col)
Vue.use(ConfigProvider)
