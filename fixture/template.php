<?php


/**
 * @param index string
 * @param id int
 * @return object
 *
 */
function help($param)
{
    return null;
}


/**
 * @apiVersion 1.0.0
 * @api {get} /mall/product/:id get detail of product
 * @apiName mall.product.get
 * @apiParam query {string{..32}} id id of product
 * @apiResponseRef /fixture/result.json
 * @apiResponse 200 {
 *  "result": true,
 *  "data": {"id": "a1", "name": "pen", "price": 101.2 }
 * } //must new line
 *
 */
function actionProductDetail($id, $name)
{
    return null;
}

/**
 * @apiVersion 1.0.0
 * @api {delete} /mall/product/:id delete product
 * @apiName mall.product.delete
 * @apiParam query {string{..32}} id id of product
 * @apiResponseRef /fixture/result.json
 * @apiResponse 201
 */
function actionDeleteProduct($id, $name)
{
    return null;
}


/**
 * @apiVersion 1.0.0
 * @api {get} /mall/products get list products
 * @apiName mall.product.list
 * @apiParam query {integer{1-}} [page=1] pageIndex
 * @apiResponseRef /fixture/result.json
 * @apiResponse 200 {
 *  "result": true,
 *  "data": [{"id": "a1", "name": "this is good", "price": 101.2, "amount": 100}]
 * } //must new line
 *
 */
function actionProductList()
{
    return null;
}


/**
 * @apiVersion 1.0.0
 * @api {post} /mall/products add new product, and return id
 * @apiName mall.product.create
 * @apiParam body {integer{100-200}} amount=123 amount of products
 * @apiParam body {string{3..50}} name name of product
 * @apiParam body {number{1-}} [price=101.2] price of product
 * @apiParam body {string{..500}} [description] description of product
 * @apiResponseRef /fixture/result.json
 * @apiResponse 200 {
 *  "result": true,
 *  "msg": "success",
 *  "data": {"id": "12312312"}
 * } //must new line
 *
 * @apiResponse 401 {
 *  "result": false,
 *  "msg": "Not Allow"
 * }//must new line
 *
 */
function actionProduct()
{
    return null;
}




/**
 * @apiVersion 1.0.0
 * @api {put} /mall/product/:id  update product
 * @apiName mall.product.update
 * @apiParam query {string{..32}} id id of product
 * @apiParam body {string{3..50}} [name] name of product
 * @apiParam body {number{1-}} [price] price of product
 * @apiParam body {integer{100-200}} [amount] amount of products
 * @apiResponseRef /fixture/result.json
 * @apiResponse 200 {
 *  "result": true,
 *  "msg": "abc",
 *  "items": ["aa", "bc"],
 *  "items": [{"id": 123, "title": "this is good"}],
 *  "data": {"id": "12312312", "abc": "very good"}
 * }//must new line
 *
 * @apiResponse 403 [{
 *  "result": false,
 *  "msg": "invalid amount"
 * }
 * ]//must new line
 *
 */
function actionUpdateProduct()
{
    return null;
}
