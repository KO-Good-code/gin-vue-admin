import service from '@/utils/request'


// @Router /json/companyList [get]
// 查询公司信息
export const getCompanyList = (params) => {
    return service({
        url: "/json/companyList",
        method: 'get',
        params
    })
}

// @Router /json/companyList [get]
// 删除公司信息
export const deleteCompanyList = (data) => {
    return service({
        url: "/json/companyList",
        method: 'delete',
        data
    })
}

// @Router /json/companyList [post]
// 更新公司信息
export const updateCompanyList = (data) => {
    return service({
        url: "/json/updateCompanyList",
        method: 'post',
        data
    })
}


// @Router /json/addCompanyList [get]
// 新增公司投标信息
export const addCompanyList = (data) => {
    return service({
        url: "/json/addCompanyList",
        method: 'post',
        data
    })
}