https://github.com/user-attachments/assets/d61e908a-590c-49b4-a663-38235462fee5

![alt text](https://github.com/fabiose81/kanban-board/blob/master/kanban-board.jpg?raw=true)

### For Golang and AWS/Cognito/Lambda
    In golang folder create a file .env and insert:

    TRUSTED_PROXIES="127.0.0.1"
    ALLOW_ORIGINS = "http://localhost:4200"
    PORT = ":9000"
    
    AWS_COGNITO_ISSUER_URL = "https://cognito-idp.%s.amazonaws.com/%s"
    AWS_REGION = {your region}   
    AWS_COGNITO_USERPOOL_ID = {your user pool id}
    AWS_COGNITO_CLIENT_ID = {your cognito client id}
    AWS_LAMBDA_FUNCTION_SAVE = {your lambda function for save}
    AWS_LAMBDA_FUNCTION_GET = {your lambda function for get}
    AWS_PROFILE = "golang" //Profile created in .aws/credentias to set aws_access_key_id and aws_secret_access_key 
    Ex: [golang]
        aws_access_key_id = {your key id}
        aws_secret_access_key = {your access key}

### Lambda code for AWS Serveless(Python) :: Save Board

    import json
    import boto3
    import uuid

    dynamodb = boto3.resource('dynamodb', region_name='us-east-1')
    table = dynamodb.Table('board')

    def lambda_handler(event, context):
        try:
            status_code = 201
            msg = "board saved successfully"

            user_id = event['userid']    
            board_id = event['boardid']
            del event['userid']
            del event['boardid']

            if board_id == "":
               board_id = str(uuid.uuid4())
               response = table.put_item(
                  Item={
                    'board_id': board_id,
                    'user_id': str(user_id),
                    'tasks': json.dumps(event)
                  }
                )
            else:
                response = table.update_item(
                  Key={
                    'board_id': str(board_id),
                    'user_id': str(user_id)                  
                  },
                  UpdateExpression="SET #tasks = :tasks",
                  ExpressionAttributeNames = {
                      '#tasks': 'tasks'
                  },
                  ExpressionAttributeValues={
                      ':tasks': json.dumps(event)
                  },
                )
                status_code = 200
                msg = "board updated successfully"
        
            return {
              'statusCode': status_code,
              'boardid' : board_id,
              'msg': msg
            }
        except Exception as e:
          print("Error saving item:", e)
          return {
            'statusCode': 400,
            'msg': e
          }


### Lambda code for AWS Serveless(Python) :: Get Board

    import boto3
    from boto3.dynamodb.conditions import Key

    dynamodb = boto3.resource('dynamodb', region_name='us-east-1')
    table = dynamodb.Table('board')

    def lambda_handler(event, context):
        try:     
            user_id = event['userid']
            response =  table.query(
              KeyConditionExpression=Key('user_id').eq(user_id)
            )
            return {
                'statusCode': 200,
                'boards': response.get('Items', [])
            }
        except Exception as e:
            print('Error querying DynamoDB:', e)
            return {
                'statusCode': 400,
                'body': e
            }
