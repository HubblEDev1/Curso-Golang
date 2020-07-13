#include <iostream>
#include "operadores.h"
using namespace std;

operaciones::operaciones(int num1,int num2){
    a=num1;
    b=num2;
}
int operaciones::suma(){
    return a+b;
}

int operaciones::resta(){
    return a-b;
}

int operaciones::mult(){
    return a*b;
}