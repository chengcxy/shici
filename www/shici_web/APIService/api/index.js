import request from '../http'


export function getNavList() {
    return request({
        url: '/api/shici/v1/dynastys/',
        method: 'get'
    })
}


export function getDynastyPoeters(params) {
  return request({
      url: '/api/shici/v1/dynasty/' + params.dynasty_id + "?currentPage="+params.currentPage+"&pageSize="+params.pageSize,
      method: 'get',
  })
}

export function getPoeterPoems(params) {
  return request({
      url: '/api/shici/v1/poeter/' + params.poeter_id+ "?currentPage="+params.currentPage+"&pageSize="+params.pageSize,
      method: 'get'
  })
}

export function getPoemDetail(params) {
  return request({
      url: '/api/shici/v1/poem/' + params.poem_id,
      method: 'get'
  })
}


