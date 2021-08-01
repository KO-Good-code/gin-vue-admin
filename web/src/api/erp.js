import service from '@/utils/request'

export const creatdJsonName = (data) => {
  return service({
      url: "/json/creatdJson",
      method: 'post',
      data
  })
}