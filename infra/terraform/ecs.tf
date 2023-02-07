resource "aws_ecs_task_definition" "ecs_task_definition" {
  family                   = local.app
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs.arn
  task_role_arn            = aws_iam_role.ecs_task.arn
  container_definitions    = <<CONTAINERS
[
    {
        "name": "${local.app}",
        "image": "medpeer/health_checks:latest",
        "portMapping": [
            {
                "hostPort": 8080,
                "containerPort": 8080,
            }
        ],
        "logConfiguration": {
            "logDriver": "awslogs",
            "options": {
                "awslogs-group": "${aws_cloudwatch_log_group.cloudwatch_log_group.name}",
                "awslogs-region": "${local.region}",
                "awslogs-stream-prefix": "${local.app}"
            }
        },
        "environment": [
            {
                "name": "NGINX_PORT",
                "value": "8080"
            },
            {
                "name": "HEALTH_CHECK_PATH",
                "value": "/health_checks"
            }
        ]
    }
]
CONTAINERS
}
