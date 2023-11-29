from django.http import HttpResponse, JsonResponse
from django.template import loader
from django.shortcuts import render
import requests
from constants import CLIENT_API_URL

def eventpage(request):
  events = requests.get(CLIENT_API_URL+"events").json()['data']
  return render(request, "event.html",{"events": events})

def eventdetailpage(request, id):
  event = requests.get(f'{CLIENT_API_URL}events/{str(id)}').json()['data'][0]
  return render(request, "eventdetail.html",{"event": event})

def clientpage(request):
  clients = requests.get(CLIENT_API_URL+"clients").json()['data']
  return render(request, "client.html", {"clients" : clients} )

def bookingpage(request):
  bookings = requests.get(CLIENT_API_URL+"bookings").json()['data']
  return render(request, "booking.html", {"bookings" : bookings} )