from django.urls import path
from . import views

urlpatterns = [
  path("clients/", views.client),
  path("bookings/", views.booking),
  path("events/", views.event),
  path("events/<str:id>", views.event_detail),
  path("seats/", views.seat),
  path("invoices/", views.pdf)
]