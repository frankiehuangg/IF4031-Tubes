from django.db import models

# Create your models here

class Client(models.Model):
  client_id = models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name="ID")
  username = models.CharField(max_length=255)
  name = models.CharField(max_length=255)
  created = models.DateTimeField(auto_created=True, auto_now_add=True)

class Booking(models.Model):

  class BookingStatus(models.TextChoices):
    COMPLETED = "COMPLETED",
    PENDING = "PENDING",
    CANCELLED = "CANCELLED",
    ERROR = "ERROR"

  booking_id = models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name="ID")
  client_id = models.IntegerField()
  event_id = models.IntegerField()
  seat_id = models.IntegerField()
  booking_status = models.CharField(max_length=255, choices= BookingStatus.choices)
  invoice_id = models.IntegerField()

  class Meta:
      unique_together = (('client_id', 'booking_id'))

'''
{
"username" : "Admin",
"name" : "Admin"
}
'''