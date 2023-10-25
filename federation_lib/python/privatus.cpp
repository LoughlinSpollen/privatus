#include <pybind11/pybind11.h>
#include <pybind11/stl.h>

#include "../cpp/include/usecase/RegisterUsecase.hpp"
#include "../cpp/include/domain/model/MemberState.hpp"
#include "../cpp/include/PrivatusLib.hpp"
namespace py = pybind11;


namespace privatus {

PYBIND11_MODULE(privatus, module) {
    module.doc() = "Privatus federated learning library";
    
    py::enum_<domain::EMemberState>(module, "MemberState")
        .value("JOINED", domain::EMemberState::JOINED)
        .value("SECEDED", domain::EMemberState::SECEDED)
        .export_values();


    py::class_<PrivatusLibrary>(module, "Privatus")
    .def(py::init([](std::string registered_model_name, FederationStateChangeObserver &federation_state_change_observer) {
        return std::unique_ptr<PrivatusLibrary>(new PrivatusLibrary(registered_model_name, federation_state_change_observer));
    }));
}

}
