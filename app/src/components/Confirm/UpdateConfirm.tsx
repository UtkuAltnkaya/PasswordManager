import { useMutation, useQueryClient } from '@tanstack/react-query';
import Modal from 'components/Modal/Modal';
import { MainContext } from 'context/provider';
import { useContext } from 'react';
import { toast } from 'react-toastify';
import { RequestToServer } from 'services/Api';
import style from './Confirm.module.scss';

const UpdateConfirm = () => {
  const { setUpdateModal, updateModal } = useContext(MainContext);

  const queryClient = useQueryClient();
  const update = useMutation(
    (item: any) => RequestToServer({ url: 'changeSitePassword', method: 'POST', data: item }),
    {
      onSuccess: () => queryClient.invalidateQueries(['user', updateModal.site]),
    },
  );

  const closeModal = () => {
    setUpdateModal({ open: false, site: '', size: 0 });
  };
  const handelClick = () =>
    update
      .mutateAsync({ site: updateModal.site, length: updateModal.size })
      .then(() => {
        toast.success('Password changed');
        closeModal();
      })
      .catch((error) => toast.error(error.message));

  return (
    <Modal onClose={closeModal}>
      <div className={style.main}>
        <div className={style.text}>
          <div>Would you like to change your old password</div>
          <small>After change old password will not be accessed</small>
        </div>
        <div className={style.buttons}>
          <button onClick={handelClick}>Yes</button>
          <button onClick={closeModal}>No</button>
        </div>
      </div>
    </Modal>
  );
};

export default UpdateConfirm;
