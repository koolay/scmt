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
 *  "extra": ["aa", "bc"],
 *  "complex": [{"id": 123, "title": "this is good"}],
 *  "data": {"id": "12312312", "abc": "very good"}
 * } //must new line
 *
 * @apiResponse 401 {
 *  "result": false,
 *  "msg": "Not Allow"
 * }//must new line
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
 *  "items": ["aa", "bc"],
 *  "items": [{"id": 123, "title": "this is good"}],
 *  "data": {"id": "12312312", "abc": "very good"}
 * }//must new line
 *
 * @apiResponse 403 [{
 *  "result": true,
 *  "msg": "not a object",
 *  "items": [{"id": 123, "title": "this is good"}]
 * }
 * ]//must new line
 *
 */
function actionUser()
{
    return null;
}


/**
 *
 * @apiVersion 1.0.1
 * @api {post} /content/:id update
 * @apiName update user
 * @apiParam id
 * @apiParam name
 * @apiParam { integer } age=18
 * @apiResponse 201
 *
 * @apiResponse 401 {
 *  "result": false,
 *  "msg": "NotAuthorization"
 * }//must new line
 */
function actionNoContent($param)
{
    if ($param == null) {
        return "aa";
    } else {
        return "111";
    }
    return null;
}

