package getngrokurl_test

import (
  "testing"
  "errors"
)

/*
**Caso de teste 1**: curlPath inválido
- Entrada: curlPath = ""
- Saída esperada: erro retornado 

**Caso de teste 2**: Comando curl falha
- Entrada: curlPath válido, mas comando curl retorna erro
- Saída esperada: erro retornado

**Caso de teste 3**: Saída do comando curl não contém URL
- Entrada: comando curl executado com sucesso, mas saída não contém URL esperada
- Saída esperada: string vazia retornada

**Caso de teste 4**: Saída do comando curl contém URL
- Entrada: comando curl executado com sucesso e saída contém URL esperada
- Saída esperada: URL extraída e retornada
*/

func TestGetngrokURL(t *testing.T) {

  // Caso 1
  _, err := GetngrokURL("")
  if err == nil {
    t.Error("Devia retornar erro para curlPath inválido")
  }

  // Caso 2
  _, err := GetngrokURL("/bin/curl_invalido") 
  if err == nil {
    t.Error("Devia retornar erro para comando curl inválido")
  }

  // Caso 3
  output := `{"tunnels": []}`
  url, _ := processRegexp(output)
  if url != "" {
    t.Error("Devia retornar string vazia para saída sem URL") 
  }

  // Caso 4
  output := `{"tunnels":[{"public_url":"https://testurl.ngrok.io"}]}` 
  url, _ := processRegexp(output)
  if url != "https://testurl.ngrok.io" {
    t.Error("Devia retornar URL extraída")
  }

}
```
