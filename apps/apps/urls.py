from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path("", include("typer.urls")),
    path("admin/", admin.site.urls),
]
