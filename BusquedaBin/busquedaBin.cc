#include "busquedaBin.h"
#include <iostream>
#include <cstdlib>
#include <ctime>
#include <algorithm>

busquedaBin::busquedaBin(int t){
    tamanio=(t>0?t:10);
    srand(time(0));
    
    for(int i=0;i<tamanio;i++){
        datos.push_back(10+rand()%90);
    }

    sort(datos.begin(),datos.end());
}
int busquedaBin::busqueda(int elemento){
    int inf=0;
    int sup=tamanio-1;
    int medio=(inf+sup+1)/2;
    int ubicacion=-1;

    do{
        mostrarSubelem(inf,sup);

        for(int i=0;i<medio;i++){
            cout<<" ";
        }
        cout<<"*";

        if(elemento==datos[medio]){
            ubicacion=medio;
        }else if(elemento<datos[medio]){
            sup=medio-1;
        }else{
            inf=medio+1;
        }
        medio=(inf+sup+1)/2;
    }while((inf<=sup)&&(ubicacion==-1));

    return ubicacion;
}

void busquedaBin::mostrar(){
    mostrarSubelem(0,tamanio-1);
}

void busquedaBin::mostrarSubelem(int inf, int sup){
    for(int i=0;i<inf;i++){
        cout<<"   ";
    }
    for(int i=inf;i<=sup;i++){
        cout<<datos[i]<<" ";
    }
    cout<<endl;
}