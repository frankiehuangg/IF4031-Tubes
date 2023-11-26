from rest_framework import serializers
from base.models import Client, Booking

class ClientSerializer(serializers.ModelSerializer):
  class Meta:
    model = Client
    fields = "__all__"


class BookingSerializer(serializers.ModelSerializer):
  class Meta:
    model = Booking
    fields = "__all__"