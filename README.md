https://github.com/user-attachments/assets/d61e908a-590c-49b4-a663-38235462fee5

![alt text](https://github.com/fabiose81/kanban-board/blob/master/kanban-board.jpg?raw=true)


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
