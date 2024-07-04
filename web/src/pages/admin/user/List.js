import React from 'react'
import { Empty } from 'antd'

function List() {
  return (
    <>
    <div>
      <h1>用户列表</h1>
    </div>
    <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
    </>
  )
}

export default List
