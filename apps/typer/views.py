from django.http import HttpResponse, HttpResponseRedirect
from django.shortcuts import render
from django.urls import reverse
from django.views.generic import ListView
from django_tables2 import SingleTableView

from .models import Record


# def index(request, template_name="typer/index.html"):
#     records = Record.objects.order_by("-speed")
#     print(records)
#     return render(request, "typer/index.html", {"records": records})

import django_tables2 as tables


class RecordTable(tables.Table):
    class Meta:
        model = Record
        template_name = "django_tables2/bootstrap.html"
        fields = ("user", "speed", "date")
        order_by = ("-speed",)


class RecordList(SingleTableView):
    model = Record
    template_name = "typer/index.html"
    table_class = RecordTable


def submit(request):
    speed = request.POST["speed"]
    record = Record(user=request.user, speed=speed)
    record.save()
    return HttpResponseRedirect(reverse("index"))
