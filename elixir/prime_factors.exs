# defmodule PrimeFactors do
#   @doc """
#   Compute the prime factors for 'number'.
#   The prime factors are prime numbers that when multiplied give the desired
#   number.
#   The prime factors of 'number' will be ordered lowest to highest.
#   """

#   @spec factors_for(pos_integer) :: [pos_integer]

#   def factors_for(number) do
#     Prime.make_primes_list(:math.sqrt(number))

#     do_factors(number, Prime.next_primes(), [])
#   end

#   defp do_factors(1, _, factors) do
#     Prime.shutdown_primes()

#     Enum.reverse(factors)
#   end

#   defp do_factors(remaining, :no_more_primes, factors) do
#     Prime.shutdown_primes()
#     IO.inspect(remaining)
#     [remaining | factors] |> Enum.reverse()
#   end

#   defp do_factors(number, divisor, factors) do
#     if rem(number, divisor) == 0 do
#       do_factors(
#         div(number, divisor),
#         divisor,
#         [divisor | factors]
#       )
#     else
#       do_factors(
#         number,
#         Prime.next_primes(),
#         factors
#       )
#     end
#   end
# end

# defmodule Prime do
#   use Agent

#   def make_primes_list(max) do
#     Agent.start_link(fn -> primes(max) end, name: __MODULE__)
#   end

#   def next_primes do
#     Agent.get_and_update(__MODULE__, fn
#       [] -> {:no_more_primes, []}
#       [first | rest] -> {first, rest}
#     end)
#   end

#   def shutdown_primes do
#     Agent.stop(__MODULE__)
#   end

#   def primes(max) do
#     IO.inspect(max)
#     Stream.iterate(2, &(&1 + 1))
#     |> Stream.filter(&prime?/1)
#     |> Stream.take_while(&(&1 <= max))
#     |> Enum.to_list()
#   end

#   def prime?(2), do: true
#   def prime?(n) when n < 2 or rem(n, 2) == 0, do: false
#   def prime?(n), do: prime?(n, 3)
#   defp prime?(n, k) when n < k * k, do: true
#   defp prime?(n, k) when rem(n, k) == 0, do: false
#   defp prime?(n, k), do: prime?(n, k + 2)
# end

# PrimeFactors.factors_for(93_819_012_551) |> IO.inspect()


defmodule PrimeFactors do
  @doc """
  Compute the prime factors for 'number'.

  The prime factors are prime numbers that when multiplied give the desired
  number.

  The prime factors of 'number' will be ordered lowest to highest.
  """
  @spec factors_for(pos_integer) :: [pos_integer]
  def factors_for(1), do: []
  def factors_for(number) do
    get_factors(number, [])
  end
  
  defp get_factors(number, acc) do
    get_factors(number, acc, 2)
  end

  defp get_factors(remaining, acc, factor) when factor * factor > remaining do
    [remaining | acc]
    |> Enum.reverse()
  end

  defp get_factors(number, acc, factor) when rem(number, factor) == 0 do
    get_factors(div(number, factor), [factor | acc], factor)
  end

  defp get_factors(number, acc, factor) do
    get_factors(number, acc, factor + 1)
  end
end

PrimeFactors.factors_for(93_819_012_551) |> IO.inspect()
