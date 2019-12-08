import React, { useState, useEffect } from 'react';
import { Animated, Easing } from 'react-native';
import LottieView from 'lottie-react-native';

const LottieLoader = ({
  animateFilePath
}) => {
  const [progress, _setProgress] = useState(() => new Animated.Value(0));

  useEffect(() => {
    Animated.timing(progress, {
      toValue: 1,
      duration: 5000,
      easing: Easing.linear
    }).start();
  }, []);

  return (
    <LottieView
      source={animateFilePath}
      progress={progress}
    />
  );
};

export default LottieLoader;
