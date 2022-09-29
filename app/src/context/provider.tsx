import { useContext, useState, FC, ReactElement, createContext } from 'react';

interface UpdateModal {
  open: boolean;
  site: string;
  size: number;
}

interface DeleteModal {
  open: boolean;
  site: string;
}

interface ContextInterface {
  user: boolean;
  updateModal: UpdateModal;
  deleteModal: DeleteModal;
  setUser?: (item: boolean) => void;
  setUpdateModal?: (item: UpdateModal) => void;
  setDeleteModal?: (item: DeleteModal) => void;
}

const defaultState = {
  user: false,
  updateModal: {} as UpdateModal,
  deleteModal: {} as DeleteModal,
};

const MainContext = createContext<ContextInterface>(defaultState);

const MainProvider: FC<{ children: ReactElement }> = ({ children }) => {
  const [user, setUser] = useState<boolean>(false);
  const [updateModal, setUpdateModal] = useState<UpdateModal>({ open: false, site: '', size: 0 });
  const [deleteModal, setDeleteModal] = useState<DeleteModal>({ open: false, site: '' });

  const data = {
    user,
    updateModal,
    deleteModal,
    setUser,
    setUpdateModal,
    setDeleteModal,
  };

  return <MainContext.Provider value={data}>{children}</MainContext.Provider>;
};

export { useContext, MainProvider, MainContext };
