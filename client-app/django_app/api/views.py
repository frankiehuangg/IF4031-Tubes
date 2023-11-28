from rest_framework.response import Response
from rest_framework import status
from rest_framework.decorators import api_view
from base.models import Client, Booking
from .serializer import ClientSerializer, BookingSerializer

@api_view(["GET", "POST" ])
def client(request):
  if (request.method == "GET"):
    clients = Client.objects.all()
    serializer = ClientSerializer(clients, many=True)
    return Response(serializer.data)

  elif (request.method == "POST"):
    serializer = ClientSerializer(data=request.data)
    if (serializer.is_valid()):
      serializer.save()
      return Response(serializer.data, status=status.HTTP_201_CREATED)
    else :
      return Response(serializer.data, status=status.HTTP_422_UNPROCESSABLE_ENTITY)
    
    

@api_view(['GET', "POST", "PATCH"])
def booking(request):
  if (request.method == "GET"):
    bookings = Booking.objects.all()
    serializer = BookingSerializer(bookings, many=True)
    return Response(serializer.data)

  elif (request.method == "POST"):
    serializer = BookingSerializer(data=request.data)
    if (serializer.is_valid()):
      serializer.save()
      return Response(serializer.data, status=status.HTTP_201_CREATED)
    else:
      return Response(serializer.data, status=status.HTTP_422_UNPROCESSABLE_ENTITY)
  
  elif (request.method == "PATCH"):
    serializer = BookingSerializer(data=request.data)
    if (serializer.is_valid()):
      serializer.save()
      return Response(serializer.data)
    else:
      return Response(serializer.data, status=status.HTTP_422_UNPROCESSABLE_ENTITY)
  
