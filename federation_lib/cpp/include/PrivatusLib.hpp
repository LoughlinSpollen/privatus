#ifndef PRIVATUS_FEDERATION_LIBRARY_H
#define PRIVATUS_FEDERATION_LIBRARY_H

// Domain layer
#include "FederationConfig.hpp"
// Infrastructure layer
#include "FederationClient.hpp"
// Usecase layer
#include "RegisterUsecase.hpp"

namespace privatus {

class PrivatusLibrary {
public:
    /// Constructor
    /// @param mlModel the MLModel ID that identifies the federation
    /// @param callback the callback federation event listener
    PrivatusLibrary(std::string mlModelId, const FederationStateChangeObserver &observer);

    /// Destructor
    ~PrivatusLibrary();

    void JoinFederation();

    /// Get the training configuration
    /// @return federation configuration
    domain::FederationConfig GetTrainingConfig() const;

private:
    std::shared_ptr<FederationClient> m_FederationClient;
    std::shared_ptr<usecase::RegisterUsecase> m_RegisterUsecase;

};

}
#endif