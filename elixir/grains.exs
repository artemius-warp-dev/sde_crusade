defmodule Grains do
  @doc """
  Calculate two to the power of the input minus one.
  """
  @spec square(pos_integer()) :: {:ok, pos_integer()} | {:error, String.t()}
  def square(number) when number >=1 and number <= 64 do
    {:ok, 2 ** (number - 1)}
  end

  def square(_), do: {:error, "The requested square must be between 1 and 64 (inclusive)"}

  @doc """
  Adds square of each number from 1 to 64.
  """
  @spec total :: {:ok, pos_integer()}
  def total do
    1..64
    |> Enum.map(&square/1)
    |> Enum.reduce(0, fn {:ok, v}, acc -> acc + v end)
    |> then(fn acc -> {:ok, acc} end)
  end
end
