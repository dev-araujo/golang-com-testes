package main

import "testing"

func TestHello(t *testing.T) {
	checkMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Recebido => '%s', Esperado => '%s'", got, want)
		}
	}

	t.Run("Diz 'Hello' para as pessoas", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		if got != want {
			checkMessage(t, got, want)
		}
	})

	t.Run("Diz 'Hello, world', quando o argumento passado é uma string vazia", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		if got != want {
			checkMessage(t, got, want)
		}
	})

	t.Run("Diz 'Hola' quando o idioma é definido como espanhol", func(t *testing.T) {
		got := Hello("Elodie", "espanhol")
		want := "Hola, Elodie"

		if got != want {
			checkMessage(t, got, want)
		}
	})

	t.Run("diz 'Bonjour' quando o idioma é definido como francês", func(t *testing.T) {
		got := Hello("Elodie", "francês")
		want := "Bonjour, Elodie"

		if got != want {
			checkMessage(t, got, want)
		}
	})

}
