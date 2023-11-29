from rest_framework.response import Response
from rest_framework import status
from rest_framework.decorators import api_view
from base.models import Client, Booking
from .serializer import ClientSerializer, BookingSerializer
from django.http import JsonResponse
import requests
from constants import TICKET_API_URL
import json

@api_view(["GET", "POST" ])
def client(request):
  if (request.method == "GET"):
    clients = Client.objects.all()
    serializer = ClientSerializer(clients, many=True)
    json_data = {"data": serializer.data, "status" : status.HTTP_200_OK, "message": "Successful select"}
    return Response(json_data)
    
  elif (request.method == "POST"):
    serializer = ClientSerializer(data=request.data)
    if (serializer.is_valid()):
      serializer.save()
      json_data = {"status": status.HTTP_200_OK, "message" : "Successful insert", "data": serializer.data}
    else :
      json_data = {"status": status.HTTP_400_BAD_REQUEST, "message" : "Failed insert", "data": None}
    return Response(json_data)
    
@api_view(['GET', "POST", "PATCH", "DELETE"])
def booking(request):
  if (request.method == "GET"):
    bookings = Booking.objects.all()
    serializer = BookingSerializer(bookings, many=True)
    json_data = {"status": status.HTTP_200_OK, "message" : "Successful select", "data": serializer.data}
    return Response(json_data)

  elif (request.method == "POST"):
    # TO DO
    serializer = BookingSerializer(data=request.data)
    if (serializer.is_valid()):
      serializer.save()
      json_data = {"status": status.HTTP_202_ACCEPTED, "message" : "Successful insert", "data": serializer.data}
    else:
      json_data = {"status": status.HTTP_400_BAD_REQUEST, "message" : "Failed insert", "data": None}
    return Response(json_data)

  elif (request.method == "PATCH"):
    # TO DO
    serializer = BookingSerializer(data=request.data)
    if (serializer.is_valid()):
      serializer.save()
      json_data ={"status": status.HTTP_202_ACCEPTED, "message" : "Successful payment", "data": serializer.data}
    else:
      json_data ={"status": status.HTTP_400_BAD_REQUEST, "message" : "Failed payment", "data": None}
    return Response(json_data)

  elif (request.method == "DELETE"):
    # TO DO
    serializer = BookingSerializer(data=request.data)
    if (serializer.is_valid()):
      serializer.save()
      json_data = {"status": status.HTTP_202_ACCEPTED, "message" : "Successful cancel", "data": serializer.data}
    else:
      json_data = {"status": status.HTTP_400_BAD_REQUEST, "message" : "Failed cancel", "data": None}
    return Response(json_data)
  
@api_view(['GET'])
def event(request):
  if (request.method == "GET"):
    events = requests.get(TICKET_API_URL+"events").json()
    json_data = {"status": status.HTTP_200_OK, "message" : "Successful select", "data": events['data']}
    return Response(json_data)

@api_view(['GET'])
def event_detail(request, id):
  if (request.method == "GET"):
    event = requests.get(TICKET_API_URL+"events/"+str(id)).json()
    json_data = {"status": status.HTTP_200_OK, "message" : "Successful select", "data": event['data']}
    return Response(json_data)

@api_view(['GET'])
def seat(request):
  if (request.method == "GET"):
    event_id = request.GET.get('event_id')
    seat_number = request.GET.get('seat_number')
    seat = requests.get(f'{TICKET_API_URL}seats?event_id={event_id}&seat_number={seat_number}').json()
    json_data = {"status": status.HTTP_200_OK, "message" : "Successful select", "data": seat['data']}
    return Response(json_data)
