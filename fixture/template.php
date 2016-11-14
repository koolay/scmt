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
 * @api {get} /user/auth/login user login
 * @apiName home page
 * @apiParam {integer{100-200}} id=123 id of user
 * @apiParam {string{..5}} [name=abc]
 * @apiParam {string{20..}} [title="good title"] title of article
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
 * @apiParam { integer } id
 * @apiParam { string } [name]
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

