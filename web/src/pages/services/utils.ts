export const isUrl = (url: string) => {
  return /^(https?:\/\/(([a-zA-Z0-9]+-?)+[a-zA-Z0-9]+\.)+[a-zA-Z]+)(:\d+)?(\/.*)?(\?.*)?(#.*)?$/.test(url);
};

export const generateQuery = (params: Object) => {
  const queryArr = [];
  for (const item of Object.entries(params)) {
    queryArr.push(item.join('='));
  }
  return `?${ queryArr.join('&') }`
};

export const serviceApiPrefix = isUrl(window.location.href) ? `${window.location.origin}/api/v1` : `${window.location.protocol}//urbs.teambition.aone.alibaba.net/api/v1`;
