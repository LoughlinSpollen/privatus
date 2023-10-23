#include <string>

#ifndef REGISTER_USECASE_H
#define REGISTER_USECASE_H

namespace usecase {

class RegisterUseCase {
public:
    /// Constructor
    /// @param mlModel the MLModel ID that identifies the federation
    /// @param callback the callback federation event listener
    RegisterUseCase(std::string mlModelId, const function &callback);

    /// Get the training configuration
    /// @return federation configuration
    FederationConfig getTrainingConfig() const;

private:
    std::string m_MlModel;
    function m_Callback;
};

}

#endif
