from django.http import HttpResponse, HttpResponseRedirect
from django.shortcuts import render
from django.urls import reverse
from django.utils import timezone

from .models import Record


def index(request, template_name="typer/index.html"):
    records = Record.objects.order_by("-speed")
    return render(request, "typer/index.html", {"records": records})


def submit(request):
    speed = request.POST["speed"]
    record = Record(speed=speed, date=timezone.now())
    record.save()
    return HttpResponseRedirect(reverse("index"))
