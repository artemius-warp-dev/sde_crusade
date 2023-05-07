defmodule Meetup do
  @moduledoc """
  Calculate meetup dates.
  """

  @teenth 13..19
  @first 1
  @second 2
  @third 3
  @fourth 4

  @type weekday ::
          :monday
          | :tuesday
          | :wednesday
          | :thursday
          | :friday
          | :saturday
          | :sunday

  @type schedule :: :first | :second | :third | :fourth | :last | :teenth

  @doc """
  Calculate a meetup date.

  The schedule is in which week (1..4, last or "teenth") the meetup date should
  fall.
  """
  @spec meetup(pos_integer, pos_integer, weekday, schedule) :: Date.t()
  def meetup(year, month, weekday, schedule) do
    start_date = Date.new!(year, month, 1)
    end_date = Date.new!(year, month, Date.days_in_month(start_date))
    range = Date.range(start_date, end_date)

    range
    |> get_date(weekday, schedule)
  end

  defp get_date(range, weekday, :teenth) do
    range
    |> Enum.find(fn date ->
      date |> Date.day_of_week(weekday) == 1 and date.day in @teenth
    end)
  end

  defp get_date(range, weekday, :last) do
    range
    |> Enum.filter(fn date ->
      date |> Date.day_of_week(weekday) == 1
    end)
    |> List.last()
  end

  defp get_date(range, weekday, final_count) when is_integer(final_count) do
    range
    |> Enum.reduce_while({nil, 1, final_count}, fn date, {_, c, fc} ->
      cond do
        date |> Date.day_of_week(weekday) == 1 and c == fc ->
          {:halt, date}

        date |> Date.day_of_week(weekday) == 1 ->
          {:cont, {nil, c + 1, fc}}

        true ->
          {:cont, {nil, c, fc}}
      end
    end)
  end

  defp get_date(range, weekday, schedule) when is_atom(schedule) do
    case schedule do
      :first -> get_date(range, weekday, @first)
      :second -> get_date(range, weekday, @second)
      :third -> get_date(range, weekday, @third)
      :fourth -> get_date(range, weekday, @fourth)
    end
  end
end

Meetup.meetup(2013, 3, :monday, :last) |> IO.inspect()
# ~D[2013-05-13]
# ~D[2013-03-25]
