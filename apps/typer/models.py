from django.db import models


class Record(models.Model):
    # name = models.CharField(max_length=200)
    speed = models.IntegerField()
    date = models.DateTimeField("date published")


# class User(models.Model):
