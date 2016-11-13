<?php

/**
 * @apiVersion 1.0.0
 * @api {get} /user/auth/login user login
 * @apiName home page
 * @apiParam { int } id required default
 * @apiParam { name } string required
 * @apiResponseRef /fixture/result.json
 * @apiResponse 200 {
 *  "result": true,
 *  "msg": "abc",
 *  "items": ["aa", "bc"]
 *  "items": [{"id": 123, "title": "this is good"}]
 *  "data": {"id": "12312312", "abc": "very good"}
 * }
 *
 * @apiResponse 403 {
 *  "result": true,
 *  "msg": "abc",
 *  "items": ["aa", "bc"]
 *  "items": [{"id": 123, "title": "this is good"}]
 *  "data": {"id": "12312312", "abc": "very good"}
 * }
 *
 */
function actionIndex($id, $name)
{
    return null;
}



/**
 * @apiVersion 1.0.0
 * @api {get} /user/:id  get user
 * @apiName get username
 * @apiParam { int } id required default
 * @apiParam { name } string required
 * @apiResponseRef /fixture/result.json
 * @apiResponse 200 {
 *  "result": true,
 *  "msg": "abc",
 *  "items": ["aa", "bc"]
 *  "items": [{"id": 123, "title": "this is good"}]
 *  "data": {"id": "12312312", "abc": "very good"}
 * }
 *
 * @apiResponse 403 {
 *  "result": true,
 *  "msg": "abc",
 *  "items": ["aa", "bc"]
 *  "items": [{"id": 123, "title": "this is good"}]
 *  "data": {"id": "12312312", "abc": "very good"}
 * }
 *
 */
function actionUser()
{
    return null;
}

