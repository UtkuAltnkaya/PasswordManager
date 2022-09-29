import { ReactNode, useRef } from 'react';
import { motion } from 'framer-motion';

import style from './Modal.module.scss';
import useOnClickOutside from 'hooks/useOnClickOutside';
import ReactDOM from 'react-dom';

type Props = {
  children: ReactNode;
  onClose: Function;
};

const Modal = ({ children, onClose }: Props) => {
  const ref = useRef(null);
  useOnClickOutside(ref, onClose);

  return ReactDOM.createPortal(
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      exit={{ opacity: 0 }}
      transition={{ duration: 0.3 }}
      className={style.main}
    >
      <div ref={ref} className={style.item}>
        {children}
      </div>
    </motion.div>,
    document.querySelector('#root'),
  );
};

export default Modal;
