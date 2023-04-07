defmodule LibraryFees do
  def datetime_from_string(string) do
    NaiveDateTime.from_iso8601!(string)
  end

  def before_noon?(datetime) do
    noon = ~T[12:00:00]
    time = NaiveDateTime.to_time(datetime)
    Time.compare(noon, time) == :gt
  end

  def return_date(checkout_datetime) do
    before_noon?(checkout_datetime)
    |> get_seconds()
    |> then(&NaiveDateTime.add(checkout_datetime, &1))
    |> NaiveDateTime.to_date()
  end

  defp get_seconds(before_noon?) when before_noon?, do: 24 * 60 * 60 * 28

  defp get_seconds(_before_noon?), do: 24 * 60 * 60 * 29 

  def days_late(planned_return_date, actual_return_datetime) do
    NaiveDateTime.to_date(actual_return_datetime)
    |> Date.diff(planned_return_date)
    |> days_late()
  end

  defp days_late(diff) when diff <= 0, do: 0
  defp days_late(diff), do: diff

  def monday?(datetime) do
    NaiveDateTime.to_date(datetime)
    |> Date.day_of_week() == 1
  end

  def calculate_late_fee(checkout, return, rate) do
    ndate_return =
      return
      |> datetime_from_string()

    ndate_checkout =
      checkout
      |> datetime_from_string()
      |> return_date()

    days_late(ndate_checkout, ndate_return)
    |> calculate_rate(rate, monday?(ndate_return))
  end

  defp calculate_rate(days, rate, monday?) when monday?, do: Kernel.floor(days * rate * 0.5)
  defp calculate_rate(days, rate, _), do: days * rate
end

