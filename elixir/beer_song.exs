defmodule BeerSong do
  @doc """
  Get a single verse of the beer song
  """
  @spec verse(integer) :: String.t()

  def verse(0) do
    """
    No more bottles of beer on the wall, no more bottles of beer.
    Go to the store and buy some more, 99 bottles of beer on the wall.
    """
  end

  def verse(number) do
    """             
    #{remind_bottles(number)} of beer on the wall, #{remind_bottles(number)} of beer.
    Take #{if number - 1 == 0, do: "it", else: "one"} down and pass it around, #{remind_bottles(number - 1)} of beer on the wall.
    """
  end



  defp remind_bottles(1) do
    "#{1} bottle"
  end

  defp remind_bottles(0) do
    "no more bottles"
  end

  defp remind_bottles(n) do
    "#{n} bottles"
  end

  @doc """
  Get the entire beer song for a given range of numbers of bottles.
  """
  @spec lyrics(Range.t()) :: String.t()
  def lyrics(range \\ 99..0) do
    range
    |> Enum.map(&verse/1)
    |> Enum.join("\n")
  end
end
