defmodule Clock do
  defstruct hour: 0, minute: 0

defimpl String.Chars, for: Clock do
  def to_string(clock) do
    hour_str = "#{clock.hour}"
    minute_str = "#{clock.minute}"
    "#{String.pad_leading(hour_str, 2, "0")}:#{String.pad_leading(minute_str, 2, "0")}"
  end
end

  @doc """
  Returns a clock that can be represented as a string:

      iex> Clock.new(8, 9) |> to_string
      "08:09"
  """
  @spec new(integer, integer) :: Clock
  def new(hour, minute) when hour >= 0 and minute >= 0 do
    %Clock{hour: rem(div(minute, 60) + hour, 24), minute: rem(minute, 60)}
  end

  def new(hour, minute) when minute >= 0  do
    %Clock{hour: 24 - Kernel.abs(rem(hour, 24)) + rem(div(minute, 60), 24), minute: rem(minute, 60)}
  end

def new(hour, minute) do
    minutes = if Kernel.abs(minute) == 60, do: rem(minute, 60), else: 60 - Kernel.abs(rem(minute, 60))
    h = if Kernel.abs(minute) > 60, do: rem(Kernel.abs(div(minute, 60)), 24) + 1, else: 1
    new(hour - h, minutes)
  end



  @doc """
  Adds two clock times:

      iex> Clock.new(10, 0) |> Clock.add(3) |> to_string
      "10:03"
  """
  @spec add(Clock, integer) :: Clock
  def add(%Clock{hour: hour, minute: minute}, add_minute) do
    new(hour, minute + add_minute)
  end
end
