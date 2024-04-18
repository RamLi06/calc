package main 

type sum struct{
  operation
}
func (sum) calc() (float32, error){
    resposta := valor1 + valor2
  return resposta, nil
}
