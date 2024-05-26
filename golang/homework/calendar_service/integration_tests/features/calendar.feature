Feature: Calendar Service
  Scenario: Add an event to the calendar
   Given an empty calendar
   When I add an event with title "Meeting" at "2024-05-26 10:00"
   Then the event "Meeting" should be added to the calendar

  Scenario: Get list of events for a day
    Given a calendar with events
    When I request events for the day "2024-05-26"
    Then I should get the events:
      | title   | time              |
      | Meeting | 2024-05-26 10:00  |

  Scenario: Send notification for an event
    Given a calendar with an event
    When the event "Meeting" time is "2024-05-26 10:00"
    Then a notification should be sent for the event "Meeting"