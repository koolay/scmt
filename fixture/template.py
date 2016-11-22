#!/bin/python


def help(param):
    """
    @param index string
    @param id int
    @return object
    """

    return None

def actionProductDetail(id, name):
    """
    @apiVersion 1.0.0
    @api {get} /mall/product/:id get detail of product
    @apiGroup product
    @apiParam query {string{..32}} id id of product
    @apiResponseRef /fixture/result.json
    @apiResponse 200 {
    "result": true,
    "data": {"id": "a1", "name": "pen", "price": 101.2 }
    } //must new line
    """

    return None

def actionDeleteProduct(id):

    """
    @apiVersion 1.0.0
    @api {delete} /mall/product/:id delete product
    @apiGroup product
    @apiParam query {string{..32}} id id of product
    @apiResponseRef /fixture/result.json
    @apiResponse 201
    """

    return None


def actionProductList():

    """
    @apiVersion 1.0.0
    @api {get} /mall/products get list products
    @apiGroup product
    @apiParam query {integer{1-}} [page=1] pageIndex
    @apiResponseRef /fixture/result.json
    @apiResponse 200 {
    "result": true,
    "data": [{"id": "a1", "name": "this is good", "price": 101.2, "amount": 100}]
    } //must new line
    """

    return None


def actionProduct():

    """
    @apiVersion 1.0.0
    @api {post} /mall/products add new product, and return id
    @apiGroup product
    @apiParam formData {integer{100-200}} amount=123 amount of products
    @apiParam formData {string{3..50}} name name of product
    @apiParam formData {number{1-}} [price=101.2] price of product
    @apiParam formData {string{..500}} [description] description of product
    @apiResponseRef /fixture/result.json
    @apiResponse 200 {
    "result": true,
    "msg": "success",
    "data": {"id": "12312312"}
    } //must new line

   @apiResponse 401 {
   "result": false,
   "msg": "Not Allow"
   }//must new line
   """

    return None




def actionUpdateProduct():

    """
    @apiVersion 1.0.0
    @api {put} /mall/product/:id  update product
    @apiGroup product
    @apiParam query {string{..32}} id id of product
    @apiParam formData {string{3..50}} [name] name of product
    @apiParam formData {number{1-}} [price] price of product
    @apiParam formData {integer{100-200}} [amount] amount of products
    @apiResponseRef /fixture/result.json
    @apiResponse 200 {
    "result": true,
    "msg": "abc",
    "items": ["aa", "bc"],
    "items": [{"id": 123, "title": "this is good"}],
    "data": {"id": "12312312", "abc": "very good"}
    }//must new line

   @apiResponse 403 [{
   "result": false,
   "msg": "invalid amount"
   }
   ]//must new line
   """

    return null
