import request from '@/utils/request'
//add crops
export function createCrop(data) {
    return request({
      url: '/createCrop',
      method: 'post',
      data
    })
  }