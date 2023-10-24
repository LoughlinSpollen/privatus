#include <pybind11/pybind11.h>
#include <pybind11/stl.h>

#include "../cpp/include/usecase/RegisterUsecase.hpp"
#include "../cpp/include/domain/model/MemberState.hpp"

namespace py = pybind11;


namespace privatus {

PYBIND11_MODULE(privatus, m) {
    m.doc() = "Privatus federated learning library";
    
    py::enum_<MemberState>(module, "MemberState")
        .value("JOINED", domain::EMemberState::JOINED)
        .value("SECEDED", domain::EMemberState::SECEDED)
        .export_values();


    py::class_<PrivatusLib>(m, "Privatus")
    .def(py::init([](std::string registered_model_name, py::object federation_state_change_observer) {
        return std::unique_ptr<PrivatusLib>(new PrivatusLib(registered_model_name, federation_state_change_observer));
    }))
}

}
