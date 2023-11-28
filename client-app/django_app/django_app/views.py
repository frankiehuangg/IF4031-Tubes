from django.http import HttpResponse, JsonResponse
from django.template import loader
from django.shortcuts import render
import requests
from constants import CLIENT_API_URL

def homepage(request):
  template = loader.get_template('event.html')
  return HttpResponse(template.render())

def clientpage(request):
  clients = requests.get(CLIENT_API_URL+"clients").json()
  print(clients)
  return render(request, "client.html", {"clients" : clients} )

def bookingpage(request):
  template = loader.get_template('booking.html')
  return HttpResponse(template.render())