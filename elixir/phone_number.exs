defmodule PhoneNumber do
  @doc """
  Remove formatting from a phone number if the given number is valid. Return an error otherwise.
  """
  @spec clean(String.t()) :: {:ok, String.t()} | {:error, String.t()}
  def clean(raw) do
    raw
    |> format_str()
    |> validate()
  end

  defp format_str(str) do
    str
    |> String.trim()
    |> String.replace(~r/^(\+?)|([" ()\-."])/, "")
  end

  defp validate(str) do
    with :ok <- ensure_is_only_digits(str),
         {:ok, validated_length_str} <- validate_length(str),
         <<area::binary-size(3), exchange::binary-size(3), _subscriber::binary-size(4)>> = validated_length_str,
         :ok <- ensure_area_is_valid(area),
         :ok <- ensure_exchange_is_valid(exchange) do
      {:ok, validated_length_str}
    end
  end

  defp ensure_area_is_valid(area) do
    cond do
      String.first(area) == "0" ->
        {:error, "area code cannot start with zero"}

      String.first(area) == "1" ->
        {:error, "area code cannot start with one"}

      true ->
        :ok
    end
  end

  defp ensure_exchange_is_valid(exchange) do
    cond do
      String.first(exchange) == "0" ->
        {:error, "exchange code cannot start with zero"}

      String.first(exchange) == "1" ->
        {:error, "exchange code cannot start with one"}

      true ->
        :ok
    end
  end

  defp ensure_is_only_digits(str) do
    if String.match?(str, ~r/[^0-9]/) do
      {:error, "must contain digits only"}
    else
      :ok
    end
  end

  defp validate_length(str) do
    case String.length(str) do
      11 -> validate_country_code(str)
      10 -> {:ok, str}
      _ -> {:error, "incorrect number of digits"}
    end
  end

  defp validate_country_code(<<"1", rest::binary>>), do: {:ok, rest}
  defp validate_country_code(_str), do: {:error, "11 digits must start with 1"}
end

PhoneNumber.clean("+1 (613)-995-0253") |> IO.inspect()
