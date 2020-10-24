/*
 *
 *     Copyright 2020 yunqi
 *
 *     Licensed under the Apache License, Version 2.0 (the "License");
 *     you may not use this file except in compliance with the License.
 *     You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 *     Unless required by applicable law or agreed to in writing, software
 *     distributed under the License is distributed on an "AS IS" BASIS,
 *     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *     See the License for the specific language governing permissions and
 *     limitations under the License.
 *
 */

package image

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/yunqi/flow"
	"github.com/yunqi/flow/function/ffile"
	"github.com/yunqi/flow/function/fimage"
	"image"
	"math/rand"
	"strconv"
	"testing"
)

func TestImage(t *testing.T) {
	newFlow := flow.NewFlow(10)
	pathFlow := newFlow.To(ffile.GetAllFiles(".jpg"))
	openFlow := pathFlow.To(fimage.OpenWithPath())
	h1 := openFlow.To(fimage.CropAnchor(300, 300, imaging.Center))
	h2 := h1.To(func(in flow.Data) (flow.Data, error) {
		return in, nil
	})
	h2.To(func(in flow.Data) (flow.Data, error) {
		data, err := fimage.Grayscale()(in)
		if err == nil {
			return fimage.Invert()(data)
		}
		fmt.Println(err)
		return nil, err
	})

	newFlow.Run(false)
	paths := []string{"data/", "d"}

	rand.Seed(2020)
	for _, path := range paths {
		newFlow.Feed(path, func(result flow.Data) {

			if re, ok := result.(error); ok {
				fmt.Println(path, re)
			}
			if ims, ok := result.([]image.Image); ok {
				for _, im := range ims {
					_ = imaging.Save(im, strconv.Itoa(rand.Int())+".jpg")
				}
			}
		})
	}

	newFlow.Wait()
}
