import { Request, Response } from 'express';
import { mock, Random} from 'mockjs';

const list = mock({
  'result|3-4':[{
    'id|+1': 1,
    'product_id': () => Random.word(24).toUpperCase(), 
    'name': () => Random.cword(3, 5),
    'desc': () => Random.sentence(5, 10),
    'status|1-3': 1,
    'created_at': () => Random.datetime(),
    'updated_at': () => Random.datetime(),
    'offline_at': () => Random.datetime(),
  }]
});

const productLogList = mock({
  'result|50-100':[{
    'id|+1': 1,
    'product': () => Random.word(24).toUpperCase(), 
    'name': () => Random.cword(3, 5),
    'desc': () => Random.sentence(5, 10),
    'channels|1-3': ['beta', 'dev', 'stable'],
    'clients|1-3': ['ios', 'web', 'andriod'],
    'status|1-3': 1,
    'release|+1': 100,
    'created_at': () => Random.datetime(),
    'updated_at': () => Random.datetime(),
    'offline_at': () => Random.datetime(),
    'users|1-10': [{
      'uid|+1': 1,
      'name': () => Random.word(4).toUpperCase(), 
    }],
  }]
});

const productModuleList = mock({
  'result|50-100':[{
    'id|+1': 1,
    'product': () => Random.word(24).toUpperCase(), 
    'name': () => Random.cword(3, 5),
    'desc': () => Random.sentence(5, 10),
    'channels|1-3': ['beta', 'dev', 'stable'],
    'clients|1-3': ['ios', 'web', 'andriod'],
    'status|1-3': 1,
    'release|+1': 100,
    'created_at': () => Random.datetime(),
    'updated_at': () => Random.datetime(),
    'offline_at': () => Random.datetime(),
    'users|1-10': [{
      'uid|+1': 1,
      'name': () => Random.word(4).toUpperCase(), 
    }],
  }]
});

const productInfo = mock({
  'result': {
    'id': 1,
    'name': () => Random.cword(3, 5),
    'desc': () => Random.sentence(5, 10),
    'labels|1-200': 20,
    'modules|1-200': 20,
    'setting|1-200': 20,
    'publish|2000-20000': 2010,
    'users|1-5': [{
      'uid|+1': 1,
      'name': () => Random.word(4).toUpperCase(), 
    }],
  }
});

const getProductInfo = (req: Request, res: Response, u: string) => {  
  return res.json(productInfo)
}

const getProducts = (req: Request, res: Response, u: string) => {  
  return res.json(list)
}

const getProductLogs = (req: Request, res: Response, u: string) => {  
  return res.json(productLogList)
}

const getProductModules = (req: Request, res: Response, u: string) => {  
  return res.json(productModuleList)
}

export default {
  'GET /api/products': getProducts,
  'GET /api/v1/products/:product/labels': getProductLogs,
  'GET /api/v1/products/:product/modules': getProductModules,
  'GET /api/v1/products/:product': getProductInfo,
}
