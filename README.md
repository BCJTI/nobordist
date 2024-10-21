# nobordist
Nobordist sdk

**Documentação:**
https://gist.github.com/nobordist/0bd0c3d72a0ceff87b553b7ca58ee643

Seguem os exemplos de requisição da API de pagamento da Nobordist:
São 3 requisições:

* 1) Authenticate para se logar à API 
  * Para se autenticar, chamar a requisição de autenticação e recuperar o token no campo
bearer_token. Este token deverá ser usado nas 2 outras requisições
  * A validade do token é de 24 horas

* 2) Creation para enviar as informações de pagamento à NB
  * Máximo 15 elementos por requisito
  * Dois valores são possíveis para o campo payment_type: icms ou darf
  * No caso de icms, é preciso criar um objeto por pedido
  * Os campos obrigatorios neste caso sao: master_number, tracking_number, dir_number, dir_date products_description, value,
customer_name, customer_state, customer_cpf
  * O CPF tem que ser válido e conter apenas dígitos
  * O CEP tem que ser um string de 8 dígitos
  * A data tem que ser um string no formato ‘YYYY-MM-DD’
  * O estado tem que ser um string de 2 letras
  * O value tem que ser um float de menos de 999999 R$, corresponde ao valor do imposto a ser pago
  * No caso de darf, pode se juntar vários pedidos no mesmo boleto
  * Os campos obrigatórios neste caso sao: value, barcode, courrier_name, courrier_cnpj
  * O value tem que ser um float de menos de 999999 R$, corresponde ao valor do imposto a ser pago
  * O barcode corresponde ao código bar do boleto
  * No caso de sucesso, um campo adicional reference_number ser retomado, este número
será necessário para consultar os status posteriormente.
  * No caso de erro, um campo adicional errors irá mostrar os erros ocorridos

* 3) Consult para consultar, em qualquer momento, o status de pagamento dos pedidos
  * Máximo 100 números de referência por requisição
  * Se um ou várias referencias nao são encontradas no sistema, nenhum erro irá aparecer.
Portanto, aparecerão na resposta apenas os pedidos encontrados
  * Os status possíveis são: processing, processed, error
  * No caso de erro, o campo messages  será preenchido com o motivo
  * No caso de sucesso, os campos authentication, barcode e payment_date serao preenchidos (no caso de DARF, já é preenchido na criação)
  * Authentication contém o código de autenticação bancária
  * Payment date contém a data que o pagamento foi processado

**Os accessos**

**Homologaçao:**
Field|Value
--------- | ------
Url|hmlapi.nobordist.com
User|phx@phx.com
Password|OK278s.LLA23h

**Produçao:**
Field|Value
--------- | ------
Url|api.nobordist.com
User|phx@phx.com
Password|(ESTÁ NO EMAIL)

