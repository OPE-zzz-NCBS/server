from locust import HttpLocust, TaskSet

def users(l):
	l.client.get("/api/users")

class UserBehavior(TaskSet):
	tasks = {users:1}

class WebsiteUser(HttpLocust):
	task_set = UserBehavior
	min_wiat = 5000
	max_wait = 9000
