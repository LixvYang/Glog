import Vue from 'vue'
import { Button, FormModel, Input, Icon, message } from 'ant-design-vue'

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
