from django.db import models

# Create your models here

class Client(models.Model):
  username = models.CharField(max_length=255)
  name = models.CharField(max_length=255)
  created = models.DateTimeField(auto_created=True, auto_now_add=True)

class Booking(models.Model):

  class BookingStatus(models.TextChoices):
    COMPLETED = "COMPLETED",
    PENDING = "PENDING",
    CANCELLED = "CANCELLED",
    ERROR = "ERROR"

  client_id = models.IntegerField()
  booking_id = models.IntegerField(auto_created=True)
  event_id = models.IntegerField()
  event_name = models.CharField(max_length=255)
  seat_number = models.IntegerField()
  booking_status = models.CharField(max_length=255, choices= BookingStatus.choices)

  class Meta:
      unique_together = (('client_id', 'booking_id'))