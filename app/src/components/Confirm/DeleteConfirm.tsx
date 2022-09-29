import { useMutation, useQueryClient } from '@tanstack/react-query';
import Modal from 'components/Modal/Modal';
import { MainContext } from 'context/provider';
import { useContext } from 'react';
import { toast } from 'react-toastify';
import { RequestToServer } from 'services/Api';
import style from './Confirm.module.scss';

const DeleteConfirm = () => {
  const { deleteModal, setDeleteModal } = useContext(MainContext);

  const queryClient = useQueryClient();
  const deleteMut = useMutation(
    (item: any) => RequestToServer({ url: 'deleteSitePassword', method: 'DELETE', params: item }),
    {
      onSuccess: () => queryClient.invalidateQueries(['user']),
    },
  );

  const closeModal = () => {
    setDeleteModal({ open: false, site: '' });
  };

  const handelClick = () =>
    deleteMut
      .mutateAsync({ site: deleteModal.site })
      .then((item) => {
        toast.success(item.message);
        closeModal();
      })
      .catch((error) => toast.error(error.message));

  return (
    <Modal onClose={closeModal}>
      <div className={style.main}>
        <div className={style.text}>
          <div>Would you like to delete your</div>
          <small>Password will not be accessed</small>
        </div>
        <div className={style.buttons}>
          <button onClick={handelClick}>Yes</button>
          <button onClick={closeModal}>No</button>
        </div>
      </div>
    </Modal>
  );
};

export default DeleteConfirm;
