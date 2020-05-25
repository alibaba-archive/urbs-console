export const generateQuery = (params: Object) => {
  const queryArr = [];
  for (const item of Object.entries(params)) {
    queryArr.push(item.join('='));
  }
  return `?${queryArr.join('&')}`
};
// 本地开发需修改 API_ORIGIN
const API_ORIGIN = window.location.origin;
export const serviceApiPrefix = `${API_ORIGIN}${window.location.pathname.includes('urbs') ? '/urbs' : ''}/api/v1`;
