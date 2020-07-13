#include <vector>
using namespace std;
class busquedaBin{
    public:
        busquedaBin(int);
        int busqueda(int);
        void mostrar();
    private: 
        int tamanio;
        vector<int> datos;
        void mostrarSubelem(int,int);
};