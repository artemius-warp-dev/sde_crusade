defmodule Gigasecond do
  @doc """
  Calculate a date one billion seconds after an input date.
  """
  @spec from({{pos_integer, pos_integer, pos_integer}, {pos_integer, pos_integer, pos_integer}}) ::
          {{pos_integer, pos_integer, pos_integer}, {pos_integer, pos_integer, pos_integer}}
  def from({{year, month, day}, {hours, minutes, seconds}} = erl_date) do
    NaiveDateTime.from_erl!(erl_date)
    |> NaiveDateTime.to_gregorian_seconds()
    |> then(fn {seconds, _} ->  seconds + 1_000000000 end)
    |> NaiveDateTime.from_gregorian_seconds()
    |> NaiveDateTime.to_erl()
  end
end
